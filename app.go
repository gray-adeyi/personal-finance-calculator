package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const TaxRate = 0.3
const OwnerRate = 0.55
const ProfitRate = 0.05
const OpExRate = 0.1

type FinanceResult struct {
	Taxes  float32 `json:"taxes"`
	Owner  float32 `json:"owner"`
	Profit float32 `json:"profit"`
	OpEx   float32 `json:"opEx"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CalculatePersonalFinance calculates and returns the finance result to the frontend.
//
// This is a simple calculation that can be executed in the frontend, but this is just a toy
// project that shows how communications between the frontend and backend happens in Wails
func (a *App) CalculatePersonalFinance(income float32) (*FinanceResult, error) {
	if income <= 0.0 {
		return nil, errors.New(fmt.Sprintf("The provided income of %.2f is invalid", income))
	}
	return &FinanceResult{
		Taxes:  TaxRate * income,
		Owner:  OwnerRate * income,
		Profit: ProfitRate * income,
		OpEx:   OpExRate * income,
	}, nil
}

func (a *App) ShowErrorDialog(message string) error {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Error",
		Message: message,
	})
	return err
}
