import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
)

func writeWithTransaction(w io.Writer, db string) error {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		_, err := txn.Update(ctx, "Performances", []string{"SingerId", "VenueId", "Revenue"}, []interface{}{1, 4, 11000})
		if err != nil {
			return err
		}
		_, err = txn.Update(ctx, "Performances", []string{"SingerId", "VenueId", "Revenue"}, []interface{}{2, 19, 15000})
		if err != nil {
			return err
		}
		fmt.Fprintln(w, "Successfully updated Performances table")
		return nil
	})
	return err
}
  
