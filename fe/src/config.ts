const config = {
  FRONTEND_URL: process.env.REACT_APP_FRONTEND_URL,
  BACKEND_URL: process.env.REACT_APP_BACKEND_URL,
  BACKEBND_PORT: process.env.REACT_APP_GO_PORT,
  statusDict: {
    Pending: 1,
    InProgress: 2,
    Done: 3,
  },
};

export default config;
