## Background

This is a simplified layout for a project structure we are currently evaluating.

Our main goal is to separate implementation logic from communication logic. Put another way, we may want to use services down the line, so web frameworks, protocols like HTTP (etc) must all be abstracted away from the implementation of business logic.

For this we have the following concepts *boundaries* - interfaces wrapping business logic. They represent "bounded contexts" (SOA speech) in our app. You might also call them 'concerns', 'components'. There are many synonyms.

Each boundary is implemented in its own package, which has the following structure:

```
houses/
    context/
    transport/
    boundary.go
```


The *boundary.go* file holds the implementation of the business rules for the boundary. In reality, it should probably just make calls into the context.

The *context* folder holds the structure of code that satisfies pure business logic. In our main application it will hold folders like usecase, domain, repository, etc..

The *transport* folder holds everything to do with the transport protocol and the format of the object. In our main application it holds folders like controllers, views, etc..

At the top of everything is an ExampleApplication object holding the boundaries together. It also ensures that boundaries that are required to speak to each other have each other as dependencies.

For example, the `House` boundary needs the `User` boundary to check whether the `user` exists when trying to add a `house`.

With this separation we should be able to swap out Gin for another framework, easily move to grpc if we want to distribute the User boundary, etc., without impacting out business logic layer.


Try the following requests:

* http://localhost:8000/user/Joe
* http://localhost:8000/house/21%20Jump%20Street


## Note

In our main app we are finding that we have more factory functions than appear here as the size of the codebase grows. A major problem has been to avoid import cycles.

The purpose of this example app is to demonstrate the intention of the design, so we can think about how to achieve this level of simplicity.