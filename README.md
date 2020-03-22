# Auth0 Playground

## What I want to do

Diagram generated with [sequencediagram.org][1].

![Sequence diagram](./docs/sequence-diagram.png)

## How to start the playground

```bash
$ docker-compose up --build  # rebuild is required when updating arguments!
```

## Application + Auth0 configuration

* Client ID represents the Auth0 application identification.
* Domain is your Auth0 tenant + Auth0 URL.

![Application configuration #1](./docs/application_1.png)
![Application configuration #2](./docs/application_2.png)

## Resources

* [Authenticating Your First React App](https://auth0.com/blog/authenticating-your-first-react-app/)

[1]: https://sequencediagram.org
