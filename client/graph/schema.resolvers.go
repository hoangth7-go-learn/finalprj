package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"time"

	"github.com/ghc-golang-hoangth7/finalprj/client/graph/model"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// UpsertPlane is the resolver for the upsertPlane field.
func (r *mutationResolver) UpsertPlane(ctx context.Context, plane model.PlaneInput) (*model.PlaneID, error) {
	// Prepare the request to upsert a plane
	planeReq := &pbPlanes.Plane{
		PlaneId:     plane.PlaneID,
		PlaneNumber: plane.PlaneNumber,
		TotalSeats:  int32(plane.TotalSeats),
		Status:      plane.Status,
	}

	// Make the gRPC request to upsert a plane
	planeIDRes, err := r.PlanesService.UpsertPlane(ctx, planeReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	planeID := &model.PlaneID{
		ID: planeIDRes.Id,
	}

	return planeID, nil
}

// ChangePlaneStatus is the resolver for the changePlaneStatus field.
func (r *mutationResolver) ChangePlaneStatus(ctx context.Context, input model.PlaneStatusInput) (bool, error) {
	// Prepare the request to change the plane status
	planeStatusReq := &pbPlanes.PlaneStatusRequest{
		PlaneId: input.PlaneID,
		Status:  input.Status,
	}

	// Make the gRPC request to change the plane status
	_, err := r.PlanesService.ChangePlaneStatus(ctx, planeStatusReq)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpsertFlight is the resolver for the upsertFlight field.
func (r *mutationResolver) UpsertFlight(ctx context.Context, flight model.FlightInput) (*model.FlightID, error) {
	// Prepare the request to upsert a flight
	flightReq := &pbFlights.Flight{
		Id:                   flight.ID,
		PlaneNumber:          flight.PlaneNumber,
		DeparturePoint:       flight.DeparturePoint,
		DestinationPoint:     flight.DestinationPoint,
		DepartureTime:        &timestamp.Timestamp{Seconds: flight.DepartureTime.Unix()},
		EstimatedArrivalTime: &timestamp.Timestamp{Seconds: flight.EstimatedArrivalTime.Unix()},
		RealDepartureTime:    &timestamp.Timestamp{Seconds: flight.RealDepartureTime.Unix()},
		RealArrivalTime:      &timestamp.Timestamp{Seconds: flight.RealArrivalTime.Unix()},
		Status:               flight.Status,
		AvailableSeats:       int32(flight.AvailableSeats),
	}

	// Make the gRPC request to upsert a flight
	flightIDRes, err := r.FlightsService.UpsertFlight(ctx, flightReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	flightID := &model.FlightID{
		ID: flightIDRes.Id,
	}

	return flightID, nil
}

// ChangeFlightStatus is the resolver for the changeFlightStatus field.
func (r *mutationResolver) ChangeFlightStatus(ctx context.Context, input model.FlightStatusInput) (bool, error) {
	// Prepare the request to change the flight status
	flightStatusReq := &pbFlights.FlightStatusRequest{
		FlightId: input.FlightID,
		Status:   input.Status,
	}

	// Make the gRPC request to change the flight status
	_, err := r.FlightsService.ChangeFlightStatus(ctx, flightStatusReq)
	if err != nil {
		return false, err
	}

	return true, nil
}

// BookFlight is the resolver for the bookFlight field.
func (r *mutationResolver) BookFlight(ctx context.Context, input model.BookFlightInput) (bool, error) {
	// Prepare the request to book a flight
	bookFlightReq := &pbFlights.BookFlightRequest{
		FlightId:   input.FlightID,
		SeatNumber: int32(input.SeatNumber),
	}

	// Make the gRPC request to book a flight
	_, err := r.FlightsService.BookFlight(ctx, bookFlightReq)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetPlanesList is the resolver for the getPlanesList field.
func (r *queryResolver) GetPlanesList(ctx context.Context, plane *model.PlaneInput) (*model.PlaneList, error) {
	// Prepare the request to get planes list
	planesReq := &pbPlanes.Plane{
		PlaneId:     plane.PlaneID,
		PlaneNumber: plane.PlaneNumber,
		TotalSeats:  int32(plane.TotalSeats),
		Status:      plane.Status,
	}

	// Make the gRPC request to get planes list
	planesListRes, err := r.PlanesService.GetPlanesList(ctx, planesReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	planes := make([]*model.Plane, len(planesListRes.Planes))
	for i, plane := range planesListRes.Planes {
		planes[i] = &model.Plane{
			PlaneID:     plane.PlaneId,
			PlaneNumber: plane.PlaneNumber,
			TotalSeats:  int(plane.TotalSeats),
			Status:      plane.Status,
		}
	}

	return &model.PlaneList{Planes: planes}, nil
}

// GetPlaneByID is the resolver for the getPlaneById field.
func (r *queryResolver) GetPlaneByID(ctx context.Context, id string) (*model.Plane, error) {
	// Prepare the request to get a plane by ID
	planeIDReq := &pbPlanes.PlaneId{Id: id}

	// Make the gRPC request to get a plane by ID
	planeRes, err := r.PlanesService.GetPlaneById(ctx, planeIDReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	plane := &model.Plane{
		PlaneID:     planeRes.PlaneId,
		PlaneNumber: planeRes.PlaneNumber,
		TotalSeats:  int(planeRes.TotalSeats),
		Status:      planeRes.Status,
	}

	return plane, nil
}

// GetFlightsList is the resolver for the getFlightsList field.
func (r *queryResolver) GetFlightsList(ctx context.Context, flight *model.FlightInput) (*model.FlightList, error) {
	// Prepare the request to get flights list
	flightsReq := &pbFlights.Flight{
		Id:                   flight.ID,
		PlaneNumber:          flight.PlaneNumber,
		DeparturePoint:       flight.DeparturePoint,
		DestinationPoint:     flight.DestinationPoint,
		DepartureTime:        &timestamp.Timestamp{Seconds: flight.DepartureTime.Unix()},
		EstimatedArrivalTime: &timestamp.Timestamp{Seconds: flight.EstimatedArrivalTime.Unix()},
		RealDepartureTime:    &timestamp.Timestamp{Seconds: flight.RealDepartureTime.Unix()},
		RealArrivalTime:      &timestamp.Timestamp{Seconds: flight.RealArrivalTime.Unix()},
		Status:               flight.Status,
		AvailableSeats:       int32(flight.AvailableSeats),
	}

	// Make the gRPC request to get flights list
	flightsListRes, err := r.FlightsService.GetFlightsList(ctx, flightsReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	flights := make([]*model.Flight, len(flightsListRes.Flights))
	for i, flight := range flightsListRes.Flights {
		flights[i] = &model.Flight{
			ID:                   flight.Id,
			PlaneNumber:          flight.PlaneNumber,
			DeparturePoint:       flight.DeparturePoint,
			DestinationPoint:     flight.DestinationPoint,
			DepartureTime:        time.Unix(flight.DepartureTime.Seconds, 0),
			EstimatedArrivalTime: time.Unix(flight.EstimatedArrivalTime.Seconds, 0),
			RealDepartureTime:    time.Unix(flight.RealDepartureTime.Seconds, 0),
			RealArrivalTime:      time.Unix(flight.RealArrivalTime.Seconds, 0),
			Status:               flight.Status,
			AvailableSeats:       int(flight.AvailableSeats),
		}
	}

	return &model.FlightList{Flights: flights}, nil
}

// GetFlightByID is the resolver for the getFlightById field.
func (r *queryResolver) GetFlightByID(ctx context.Context, id string) (*model.Flight, error) {
	// Prepare the request to get a flight by ID
	flightIDReq := &pbFlights.FlightId{Id: id}

	// Make the gRPC request to get a flight by ID
	flightRes, err := r.FlightsService.GetFlightById(ctx, flightIDReq)
	if err != nil {
		return nil, err
	}

	// Convert the gRPC response to the GraphQL response type
	flight := &model.Flight{
		ID:                   flightRes.Id,
		PlaneNumber:          flightRes.PlaneNumber,
		DeparturePoint:       flightRes.DeparturePoint,
		DestinationPoint:     flightRes.DestinationPoint,
		DepartureTime:        time.Unix(flightRes.DepartureTime.Seconds, 0),
		EstimatedArrivalTime: time.Unix(flightRes.EstimatedArrivalTime.Seconds, 0),
		RealDepartureTime:    time.Unix(flightRes.RealDepartureTime.Seconds, 0),
		RealArrivalTime:      time.Unix(flightRes.RealArrivalTime.Seconds, 0),
		Status:               flightRes.Status,
		AvailableSeats:       int(flightRes.AvailableSeats),
	}

	return flight, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
