const { getMovies } = require('./db');

const resolvers = {
  Query: {
    movies: (_, { limit, rating }) => getMovies(limit, rating),
  },
};

module.exports = resolvers;
