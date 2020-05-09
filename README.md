# Go Test Helpers

Helper functions for writing unit tests in Golang.

The functions are divided into different sections based on their nature or corresponding package in Go.

I have tried to follow best practices of testing wherever possible inspired by [Mitchell Hashimoto](https://speakerdeck.com/mitchellh/advanced-testing-with-go)

Any place where closing a stream/object is deferred, the defer operation is returned as a `closer` function that can be triggered in the test when invoking the function thereby cleaning up correctly after the test is completed.
