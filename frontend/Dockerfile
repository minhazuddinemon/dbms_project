# Build Stage
FROM node:22-alpine AS builder

WORKDIR /app

# Copy package manifests first to leverage Docker layer caching
COPY package*.json ./

# Install all dependencies (including devDependencies required for build)
RUN npm ci

# Copy full frontend source code
COPY . .

# Build the SvelteKit application
RUN npm run build

# Production Stage
FROM node:22-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production

# Copy package manifests to install only production dependencies
COPY package*.json ./
RUN npm ci --omit=dev

# Copy compiled build output from the builder stage
COPY --from=builder /app/build ./build

# Expose SvelteKit default Node server port
EXPOSE 3000

# Start the SvelteKit production server
CMD ["node", "build"]
