const { people, getById } = require('./db');

const resolvers = {
  Query: {
    people: () => people,
    person: (_, { id }) => getById(id),
  },
};

module.exports = resolvers;
