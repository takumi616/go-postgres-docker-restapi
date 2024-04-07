// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"sync"
)

// Ensure, that SentenceDeleterMock does implement SentenceDeleter.
// If this is not the case, regenerate this file with moq.
var _ SentenceDeleter = &SentenceDeleterMock{}

// SentenceDeleterMock is a mock implementation of SentenceDeleter.
//
//	func TestSomethingThatUsesSentenceDeleter(t *testing.T) {
//
//		// make and configure a mocked SentenceDeleter
//		mockedSentenceDeleter := &SentenceDeleterMock{
//			DeleteSentenceFunc: func(ctx context.Context, sentenceID int) (int, error) {
//				panic("mock out the DeleteSentence method")
//			},
//		}
//
//		// use mockedSentenceDeleter in code that requires SentenceDeleter
//		// and then make assertions.
//
//	}
type SentenceDeleterMock struct {
	// DeleteSentenceFunc mocks the DeleteSentence method.
	DeleteSentenceFunc func(ctx context.Context, sentenceID int) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteSentence holds details about calls to the DeleteSentence method.
		DeleteSentence []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SentenceID is the sentenceID argument value.
			SentenceID int
		}
	}
	lockDeleteSentence sync.RWMutex
}

// DeleteSentence calls DeleteSentenceFunc.
func (mock *SentenceDeleterMock) DeleteSentence(ctx context.Context, sentenceID int) (int, error) {
	if mock.DeleteSentenceFunc == nil {
		panic("SentenceDeleterMock.DeleteSentenceFunc: method is nil but SentenceDeleter.DeleteSentence was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		SentenceID int
	}{
		Ctx:        ctx,
		SentenceID: sentenceID,
	}
	mock.lockDeleteSentence.Lock()
	mock.calls.DeleteSentence = append(mock.calls.DeleteSentence, callInfo)
	mock.lockDeleteSentence.Unlock()
	return mock.DeleteSentenceFunc(ctx, sentenceID)
}

// DeleteSentenceCalls gets all the calls that were made to DeleteSentence.
// Check the length with:
//
//	len(mockedSentenceDeleter.DeleteSentenceCalls())
func (mock *SentenceDeleterMock) DeleteSentenceCalls() []struct {
	Ctx        context.Context
	SentenceID int
} {
	var calls []struct {
		Ctx        context.Context
		SentenceID int
	}
	mock.lockDeleteSentence.RLock()
	calls = mock.calls.DeleteSentence
	mock.lockDeleteSentence.RUnlock()
	return calls
}