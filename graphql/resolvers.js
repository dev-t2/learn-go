const person = {
  name: 'developer-t2',
  age: 33,
  gender: 'male',
};

const resolvers = {
  Query: {
    person: () => person,
  },
};

module.exports = resolvers;
