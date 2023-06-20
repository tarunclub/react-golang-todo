# Use the official Node.js base image
FROM node:14-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the source code into the container
COPY . .

# Build the React application
RUN npm run build

# Expose port 5173
EXPOSE 5173

# Set the entry point for the container
CMD ["npm", "run", "dev"]
