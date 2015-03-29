## Background

This is an example of evolving a golang code structure that lets a team begin coding a monolith but transparently evolve toward a service design later.

Our main goal is to separate implementation logic from communication logic. Put another way, we may want to use services down the line, so web frameworks, protocols like HTTP (etc) must all be abstracted away from the implementation of business logic.

For this we have a *boundaries* concept - interfaces wrapping business logic. They represent "bounded contexts" (SOA speech) in our app. You might also call them 'concerns', 'components'. They might one day become distributed services, talking to each other. There are many synonyms.

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


We currently have two examples - see the `rpc` and `monolith` folders.

