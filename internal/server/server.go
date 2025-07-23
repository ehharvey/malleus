package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ehharvey/malleus/internal/config"
	"github.com/ehharvey/malleus/internal/graph"
	"github.com/ehharvey/malleus/internal/graph/resolver"
	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/repository"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
)

func StartServer() {
	config.InitializeConfig()
	config.LoadConfig()

	// if configErr != nil {
	// 	log.Fatal(fmt.Errorf("error reading config!: %w", configErr))
	// 	os.Exit(1)
	// }

	resolver := &resolver.Resolver{}

	// Database
	log.Print("initializing database")
	dbDriver := viper.GetString("db.driver")
	log.Printf("db.driver is %s", dbDriver)

	if dbDriver == "postgresql" {
		connString := strings.Join(strings.Fields(fmt.Sprintf(
			`
			user=%s
			password=%s
			host=%s
			port=%d
			dbname=%s
			pool_max_conns=%d
			`,
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.host"),
			viper.GetInt("db.port"),
			viper.GetString("db.database_name"),
			viper.GetInt("db.pool.max_conns"),
		)), " ")

		log.Printf("connString=%s", connString)

		lookup, lookupErr := net.LookupHost(viper.GetString("db.host"))

		if lookupErr != nil {
			log.Fatal(fmt.Errorf("error looking up db host!: %w", lookupErr))
		} else {
			log.Printf("db host ip is %v", lookup[:])
		}

		config, dbConfigErr := pgxpool.ParseConfig(connString)

		if dbConfigErr != nil {
			log.Fatal(fmt.Errorf("error parsing db config!: %w", dbConfigErr))
			os.Exit(2)
		}

		pool, poolErr := pgxpool.NewWithConfig(
			context.Background(),
			config,
		)

		if poolErr != nil {
			log.Fatal(fmt.Errorf("error initializing db pool!: %w", poolErr))
			os.Exit(2)
		}

		// service initialization
		inventoryService := inventory.NewService(
			repository.NewInventoryRepository(
				repository.NewInventoryQueries(pool),
			),
		)

		log.Print("initalized pool and services")

		resolver.InventoryService = *inventoryService

	} else {
		log.Fatal("dbDriver is invalid")
		os.Exit(1)
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := viper.GetString("server.port")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
