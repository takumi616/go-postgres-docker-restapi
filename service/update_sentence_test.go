package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/takumi616/go-postgres-docker-restapi/entity"
)

func TestUpdateSentence(t *testing.T) {
	//Prepare two test cases
	testCases := map[string]entity.Sentence{}

	//OK
	testCases["ok"] = entity.Sentence{
		SentenceID:   6,
		Body:         "After completing the build process, the application is packaged into a container and ready for deployment.",
		Vocabularies: pq.StringArray{"build", "deployment", "container"},
		Created:      "2024-04-06 20:16:35.47968413 +0000 UTC m=+25.323730179",
		Updated:      "2024-04-06 20:16:35.47969263 +0000 UTC m=+25.323738679",
	}

	//Data does not exist
	testCases["error"] = entity.Sentence{
		SentenceID:   5,
		Body:         "After completing the build process, the application is packaged into a container and ready for deployment.",
		Vocabularies: pq.StringArray{"application", "store", "server"},
		Created:      "2024-04-06 20:16:35.47968413 +0000 UTC m=+25.323730179",
		Updated:      "2024-04-06 20:16:35.47969263 +0000 UTC m=+25.323738679",
	}

	for name, testSentence := range testCases {
		name := name
		testSentence := testSentence
		//Execute as parallel tests
		//Run runs function as a subtest of t called name n(first parameter of Run)
		//It runs function in a separate goroutine and blocks
		//until this function returns or calls t.Parallel to become a parallel test
		t.Run(name, func(t *testing.T) {
			//Parallel signals that this test is to be run in parallel
			//with (and only with) other parallel tests
			t.Parallel()

			//SentenceDeleterMock　mocks SentenceDeleter interface
			//which is used to call service package method
			ctx := context.Background()
			moq := &SentenceUpdaterMock{}
			moq.UpdateSentenceFunc = func(ctx context.Context, sentenceID int64, body string) (int64, error) {
				if testSentence.SentenceID == sentenceID {
					testSentence.Body = body
					return testSentence.SentenceID, nil
				}
				return 0, errors.New("sql: no rows in result set")
			}

			//Call test target method using mock interface
			body := "After finishing the build process, the application is packaged into a container and ready for deployment."
			u := &UpdateSentence{Store: moq}
			updatedSentence, err := u.UpdateSentence(ctx, "6", body)
			if err != nil && err.Error() != "sql: no rows in result set" {
				t.Fatalf("Failed to get expected result: %v", err)
			}

			if name == "ok" {
				t.Logf("Successfully updated a sentence: %d", updatedSentence)
			} else {
				//Test case "error"
				t.Logf("Failed to update: %v", err)
			}

		})
	}
}
