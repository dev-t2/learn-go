const people = [
  {
    id: 0,
    name: 'jyk',
    age: 33,
    gender: 'male',
  },
  {
    id: 1,
    name: 'yjl',
    age: 34,
    gender: 'female',
  },
  {
    id: 2,
    name: 'jsl',
    age: 28,
    gender: 'male',
  },
];

exports.getById = (id) => {
  const filteredPeople = people.filter((person) => person.id === id);

  return filteredPeople[0];
};

exports.people = people;
