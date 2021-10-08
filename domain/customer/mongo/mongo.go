// Mongo - это mongo реализация репозитория Customer
package mongo

import (
	"context"
	"github.com/MaksimDzhangirov/tavern/domain/customer"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
	// customer используется для сохранения клиентов
	customer *mongo.Collection
}

// mongoCustomer - это приватный тип, который используется для хранения CustomerAggregate
// мы используем приватную структуру, чтобы избежать связывания этой реализации для MongoDB с CustomerAggregate.
// Mongo использует bson, поэтому мы добавили сюда дескрипторы
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// newFromCustomer принимает на вход агрегат и преобразует его в приватную структуру
func newFromCustomer(c customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate преобразуется в aggregate.Customer
// здесь также можно осуществить валидацию всех значений
func (m mongoCustomer) ToAggregate() customer.Customer {
	c := customer.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}

// New - создаёт новый MongoDB репозиторий
func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// находим Metabot DB
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	// Преобразуем в агрегат
	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := newFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) Update(c customer.Customer) error {
	panic("to implement")
}
