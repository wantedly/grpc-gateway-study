package server

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/wantedly/grpc-gateway-study/api"
)

// NewBookServiceServer creates a new BookServiceServer instance.
func NewBookServiceServer(
	db *sqlx.DB,
) interface {
	api_pb.BookServiceServer
	grapiserver.Server
} {
	return &bookServiceServerImpl{
		db: db,
	}
}

type bookServiceServerImpl struct {
	db *sqlx.DB
}

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *bookServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterBookServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *bookServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterBookServiceHandler(ctx, mux, conn)
}

type Book struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

func (s *bookServiceServerImpl) ListBooks(ctx context.Context, req *api_pb.ListBooksRequest) (*api_pb.ListBooksResponse, error) {
	books := []Book{}
	err := s.db.SelectContext(ctx, &books, "SELECT * FROM books ORDER BY id desc")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := &api_pb.ListBooksResponse{}
	for _, b := range books {
		resp.Books = append(resp.Books, &api_pb.Book{
			BookId: fmt.Sprint(b.ID),
			Title:  b.Title,
		})
	}
	return resp, nil
}

func (s *bookServiceServerImpl) GetBook(ctx context.Context, req *api_pb.GetBookRequest) (*api_pb.Book, error) {
	var book Book
	err := s.db.SelectContext(ctx, &book, "SELECT * FROM books WHERE id = $1", req.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api_pb.Book{BookId: fmt.Sprint(book.ID), Title: book.Title}, nil
}

func (s *bookServiceServerImpl) CreateBook(ctx context.Context, req *api_pb.CreateBookRequest) (*api_pb.Book, error) {
	b := req.GetBook()
	rows, err := s.db.NamedQueryContext(ctx, "INSERT INTO books (title) VALUES (:title) RETURNING id", &Book{Title: b.GetTitle()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var id int64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &api_pb.Book{
		BookId: fmt.Sprint(id),
		Title:  b.GetTitle(),
	}, nil
}

func (s *bookServiceServerImpl) UpdateBook(ctx context.Context, req *api_pb.UpdateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) DeleteBook(ctx context.Context, req *api_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
