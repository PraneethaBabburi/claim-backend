package repository

import (
	"claims-backend/internal/model"
	"database/sql"
)

type MySQLClaimRepository struct {
	db *sql.DB
}

func NewMySQLClaimRepository(db *sql.DB) *MySQLClaimRepository {
	return &MySQLClaimRepository{db: db}
}

func (r *MySQLClaimRepository) Save(claim model.Claim) error {
	query := `INSERT INTO claims (id, user_id, type, amount, status, created_at)
			values(?,?,?,?,?,?)`
	_, err := r.db.Exec(
		query,
		claim.ID,
		claim.UserID,
		claim.Type,
		claim.Amount,
		claim.Status,
		claim.CreatedAt,
	)
	return err
}

func (r *MySQLClaimRepository) FindByID(id string) (model.Claim, error) {
	query := `SELECT id, user_id, type, amount, status, created_at
			FROM claims WHERE id=?`
	row := r.db.QueryRow(query, id)
	var claim model.Claim
	err := row.Scan(
		&claim.ID,
		&claim.UserID,
		&claim.Type,
		&claim.Amount,
		&claim.Status,
		&claim.CreatedAt,
	)
	if err != nil {
		//log.Println("FindByID scan error:", err)
		return model.Claim{}, err
	}
	return claim, nil
}

func (r *MySQLClaimRepository) FindAll() []model.Claim {
	rows, err := r.db.Query(`SELECT id, user_id, type, amount, status, created_at
			FROM claims`)
	if err != nil {
		return []model.Claim{}
	}
	defer rows.Close()
	var claims []model.Claim
	for rows.Next() {
		var claim model.Claim
		rows.Scan(
			&claim.ID,
			&claim.UserID,
			&claim.Type,
			&claim.Amount,
			&claim.Status,
			&claim.CreatedAt,
		)
		claims = append(claims, claim)
	}
	return claims
}

func (r *MySQLClaimRepository) Update(claim model.Claim) bool {
	query := `UPDATE claims SET status=?
			WHERE id=?`
	_, err := r.db.Exec(query, claim.Status, claim.ID)
	return err == nil
}
