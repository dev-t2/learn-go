const { default: axios } = require('axios');
const asiox = require('axios');

const API_URL = 'https://yts.mx/api/v2/list_movies.json?';

exports.getMovies = async (limit, rating) => {
  const {
    data: {
      data: { movies },
    },
  } = await axios(API_URL, {
    params: {
      limit,
      minimum_rating: rating,
    },
  });

  return movies;
};
