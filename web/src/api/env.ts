let baseUrl = '';
if (process.env.NODE_ENV == 'development') {
  baseUrl = `http://127.0.0.1:3000`;
} else if (process.env.NODE_ENV == 'production') {
  baseUrl = `http://10.10.10.61:8778/`;
}

export { baseUrl };
