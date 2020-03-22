FROM node:12.6.0 AS builder
WORKDIR /app
COPY . .
RUN ["yarn", "run", "build"]

FROM node:12.6.0 AS runner
RUN ["yarn", "global", "add", "serve"]
WORKDIR /app
COPY --from=builder /app/build .
EXPOSE 3000
ENV REACT_APP_AUTH0_DOMAIN=com.domain
ENV REACT_APP_AUTH0_CLIENT_ID=1234567890abcde
ENV REACT_APP_REDIRECT_URL=http://localhost/callback
CMD ["serve", "-p", "3000", "-s", "."]
