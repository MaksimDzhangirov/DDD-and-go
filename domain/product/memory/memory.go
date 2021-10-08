// Пакет memory - это реализация в памяти интерфейса ProductRepository
package memory

import (
	"github.com/MaksimDzhangirov/tavern/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

// New - это функция-фабрика для создания нового ProductRepository
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

// GetAll возвращает все товары в виде слайса
// Да, он никогда не возвращает ошибку, но, например,
// реализация для базы данных может возвращать ошибку
func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	// Извлекаем все товары из карты
	var products []product.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID ищет товар, используя его ID
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

// Add добавит новый товар в репозиторий
func (mpr *MemoryProductRepository) Add(newprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

// Update изменит все значения в товаре, найденный по ID
func (mpr *MemoryProductRepository) Update(upprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete удалит товар из репозитория
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}