package contextgraph

import (
	"context"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/importcontextgraph"
	libmongo "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoReporter struct {
	collection libmongo.DbCollection
}

func NewMongoStatusReporter(client libmongo.DbClient) StatusReporter {
	return &mongoReporter{
		collection: client.Collection(libmongo.ImportJobMongoCollection),
	}
}

func (r *mongoReporter) ReportCreate(ctx context.Context, job *ImportJob) error {
	return r.create(ctx, job)
}

func (r *mongoReporter) ReportOngoing(ctx context.Context, job ImportJob) error {
	job.Status = statusOngoing
	return r.update(ctx, job)
}

func (r *mongoReporter) ReportDone(ctx context.Context, job ImportJob, stats importcontextgraph.Stats) error {
	job.Status = statusDone
	t := time.Time{}
	job.ExecTime = t.Add(stats.ExecTime).Format("15:04:05")
	job.Stats = stats

	return r.update(ctx, job)
}

func (r *mongoReporter) ReportError(ctx context.Context, job ImportJob, execDuration time.Duration, err error) error {
	job.Status = statusFailed
	t := time.Time{}
	job.ExecTime = t.Add(execDuration).Format("15:04:05")
	job.Info = err.Error()

	return r.update(ctx, job)
}

func (r *mongoReporter) GetStatus(ctx context.Context, id string) (ImportJob, error) {
	var status ImportJob
	res := r.collection.FindOne(ctx, bson.M{"_id": id})
	err := res.Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = ErrNotFound
		}
	} else {
		err = res.Decode(&status)
	}

	return status, err
}

func (r *mongoReporter) create(ctx context.Context, job *ImportJob) error {
	job.ID = utils.NewID()

	_, err := r.collection.InsertOne(ctx, job)
	return err
}

func (r *mongoReporter) update(ctx context.Context, job ImportJob) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": job.ID}, bson.M{"$set": job})
	return err
}
