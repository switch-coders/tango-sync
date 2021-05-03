package infrastructure

import "github.com/newrelic/go-agent/v3/newrelic"

func WrapExternalSegmentWithAlias(txn *newrelic.Transaction, url string, hostAlias string, segment func()) {
	nrSegment := newrelic.ExternalSegment{
		URL:  url,
		Host: hostAlias,
	}
	nrSegment.StartTime = txn.StartSegmentNow()
	defer nrSegment.End()
	segment()
}

func WrapDatastoreSegment(product string, operation string, txn *newrelic.Transaction, segment func()) {
	sg := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product:   newrelic.DatastoreProduct(product),
		Operation: operation,
	}
	defer sg.End()
	segment()
}
