export const load = async () => {
  const response = await fetch(`http://localhost:8080/countries?page=0`, {
    method: "GET",
    headers: {
      accept: "application/json",
    },
  });
  const res = await response.json();
  let countries = res.data;
  return {
    countries,
  };
};
