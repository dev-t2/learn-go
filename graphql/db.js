const fetch = require('node-fetch');

const API_URL = 'https://yts.mx/api/v2/list_movies.json?';

exports.getMovies = (limit, rating) => {
  let requestUrl = API_URL;

  if (limit > 0 && rating > 0) {
    requestUrl = `${requestUrl}limit=${limit}&minimum_rating=${rating}`;
  } else if (limit > 0) {
    requestUrl = `${requestUrl}limit=${limit}`;
  } else if (rating > 0) {
    requestUrl = `${requestUrl}minimum_rating=${rating}`;
  }

  return fetch(requestUrl)
    .then((res) => res.json())
    .then((json) => json.data.movies);
};
