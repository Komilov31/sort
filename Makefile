all: test

test:
	cd integration_tests && mkdir results && bash tests.sh

test_multiflag:
	cd integration_tests && mkdir results && bash test_multiflag.sh

clean:
	cd integration_tests/results && rm *