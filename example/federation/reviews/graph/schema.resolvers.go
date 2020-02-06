// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	"github.com/99designs/gqlgen/example/federation/reviews/graph/generated"
	"github.com/99designs/gqlgen/example/federation/reviews/graph/model"
)

func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	switch obj.Upc {
	case "top-1":
		return []*model.Review{{
			Body: "A highly effective form of birth control.",
		}}, nil
	case "top-2":
		return []*model.Review{{
			Body: "Fedoras are one of the most fashionable hats around and can look great with a variety of outfits.",
		}}, nil

	case "top-3":
		return []*model.Review{{
			Body: "This is the last straw. Hat you will wear. 11/10",
		}}, nil

	}
	return nil, nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	if obj.ID == "1234" {
		return []*model.Review{{
			Body: "The best. better than the rest.",
		}}, nil
	}
	return nil, nil
}

func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }
func (r *Resolver) User() generated.UserResolver       { return &userResolver{r} }

type productResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
