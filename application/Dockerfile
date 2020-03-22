FROM node:12.6.0 AS builder
WORKDIR /app
COPY . .
ARG REACT_APP_AUTH0_DOMAIN=com.domain
ARG REACT_APP_AUTH0_CLIENT_ID=1234567890abcd
ARG REACT_APP_REDIRECT_URL=http://localhost/
RUN ["yarn", "run", "build"]

FROM node:12.6.0 AS runner
RUN ["yarn", "global", "add", "serve"]
WORKDIR /app
COPY --from=builder /app/build .
EXPOSE 3000
CMD ["serve", "-p", "3000", "-s", "."]
