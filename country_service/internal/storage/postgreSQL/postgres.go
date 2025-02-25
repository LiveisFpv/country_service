package postgresql

import "github.com/jackc/pgx/v5/pgxpool"

type Queries struct {
	pool *pgxpool.Pool
}

//Constructor postgres pool
func New(pgxpool *pgxpool.Pool) *Queries {
	return &Queries{pool: pgxpool}
}

//Destructor postgres pool
func (q *Queries) Stop() {
	q.pool.Close()
}
