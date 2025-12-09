package models

import "time"

// ====================
// User
// ====================
type User struct {
	ID          uint      `gorm:"primaryKey"`
	Nome        string    `gorm:"size:100;not null"`
	Email       string    `gorm:"size:150;uniqueIndex;not null"`
	SenhaHash   string    `gorm:"size:255;not null"`
	Tipo        string    `gorm:"size:50;not null"` // cliente, admin, profissional
	DataCriacao time.Time `gorm:"autoCreateTime"`
}

// ====================
// Customer
// ====================
type Customer struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null;uniqueIndex"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CPFCNPJ  string `gorm:"size:20"`
	Endereco string `gorm:"size:255"`
}

// ====================
// Professional
// ====================
type Professional struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null;uniqueIndex"`
	User          User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CREACAUU      string `gorm:"size:50"`
	Especialidade string `gorm:"size:100"`
	Bio           string `gorm:"type:text"`
}

// ====================
// Project
// ====================
type Project struct {
	ID             uint `gorm:"primaryKey"`
	ProfessionalID uint `gorm:"not null"`
	Professional   Professional

	Titulo     string `gorm:"size:200;not null"`
	Descricao  string `gorm:"type:text"`
	AreaM2     float64
	Quartos    int
	Banheiros  int
	Vagas      int
	Pavimentos int
	Estilo     string  `gorm:"size:50"`
	Preco      float64 `gorm:"type:decimal(10,2)"`
	Status     string  `gorm:"size:50;default:'ativo'"`

	Images   []ProjectImage
	Files    []ProjectFile
	Licenses []License
	Reviews  []Review
}

// ====================
// ProjectImage
// ====================
type ProjectImage struct {
	ID        uint   `gorm:"primaryKey"`
	ProjectID uint   `gorm:"index;not null"`
	URLImagem string `gorm:"size:255;not null"`
	Ordem     int
}

// ====================
// ProjectFile
// ====================
type ProjectFile struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectID   uint   `gorm:"index;not null"`
	TipoArquivo string `gorm:"size:50"`
	URLArquivo  string `gorm:"size:255;not null"`
	TamanhoMB   float64
}

// ====================
// Order
// ====================
type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint `gorm:"index;not null"`
	Customer   Customer

	ValorTotal  float64   `gorm:"type:decimal(10,2)"`
	Status      string    `gorm:"size:30;default:'pendente'"`
	DataCriacao time.Time `gorm:"autoCreateTime"`

	Items   []OrderItem
	Payment Payment
}

// ====================
// OrderItem
// ====================
type OrderItem struct {
	ID            uint    `gorm:"primaryKey"`
	OrderID       uint    `gorm:"index;not null"`
	ProjectID     uint    `gorm:"index;not null"`
	PrecoUnitario float64 `gorm:"type:decimal(10,2)"`
}

// ====================
// Payment
// ====================
type Payment struct {
	ID            uint   `gorm:"primaryKey"`
	OrderID       uint   `gorm:"uniqueIndex"`
	Gateway       string `gorm:"size:100"`
	TransacaoID   string `gorm:"size:150"`
	Status        string `gorm:"size:50"`
	DataPagamento time.Time
}

// ====================
// License
// ====================
type License struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectID   uint   `gorm:"index;not null"`
	TipoLicenca string `gorm:"size:100"`
	Termos      string `gorm:"type:text"`
}

// ====================
// Review
// ====================
type Review struct {
	ID         uint      `gorm:"primaryKey"`
	ProjectID  uint      `gorm:"index;not null"`
	CustomerID uint      `gorm:"index;not null"`
	Nota       int       `gorm:"not null"`
	Comentario string    `gorm:"type:text"`
	Data       time.Time `gorm:"autoCreateTime"`
}
