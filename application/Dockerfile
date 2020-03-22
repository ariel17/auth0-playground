FROM node:12.6.0 AS builder
WORKDIR /app
COPY . .
RUN ["yarn", "run", "build"]

FROM node:12.6.0 AS runner
RUN ["yarn", "global", "add", "serve"]
WORKDIR /app
COPY --from=builder /app/build .
EXPOSE 3000
CMD ["serve", "-p", "3000", "-s", "."]
