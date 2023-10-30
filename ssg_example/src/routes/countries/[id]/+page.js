export const load = async (loadEvent) => {
  const { params } = loadEvent;
  const { id } = params;
  const response = await fetch(`http://localhost:8080/countries?page=${id}`, {
    method: "GET",
    headers: {
      accept: "application/json",
    },
  });
  const res = await response.json();
  let countries = res.data;
  let pages = Array(res.noOfPages)
    .fill(0)
    .map((_x, i) => i);
  return {
    countries,
    pages,
  };
};
