# Acceptance testing

- Acceptance testing should always be done from the perspective of the end user, so it should be at a higher level of abstraction than the code under test. Seperation of concerns is paramount to avoid tests breaking every time the implemtations change. the point is to seperate what the code does from how it does it.

- the tests should be run in a simulated production like enviroment.

- Dave Farley's video on this illustrates a good way to write acceptance tests by using internal/external DSLs.

- the DSL should be written according to the language of the problem domain (i.e if we want to test code that orders books from a website we should have a DSL that reflects that purpose.)

- The DSL should be simple to read and be easy to use to create new test cases.

- The DSL should be shared between the test cases.

- The next part of the test should be the drivers, this is the only part of the code that should be concerned with implementation details because its purpose is to translate the DSL to the system under test.

- to summarize, the acceptance test should be structured as such: Test case -> DSL -> Drivers -> system under test.

# ProtoBufs

- Protobufs are a binary serialisation format that allows you to define the structure of your data using a specific Interface Definition Language, and then generate code in various programming languages to serialize and deserialize data according to that definition.

- protobufs are used when you need to specify smaller, simpler protocols between client and server code, keep in mind that in most cases protobufs are only used when you are in control of both the client and server