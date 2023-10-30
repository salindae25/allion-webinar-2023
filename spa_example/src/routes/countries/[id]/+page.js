// export const ssr = false;
// export const csr = false;
// export const prerender = true;
// export const prerender = "auto";
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
/** @type {import('./$types').EntryGenerator} */
// export async function entries() {
//   const response = await fetch(`http://localhost:8080/countries/0/page`, {
//     method: "GET",
//     headers: {
//       accept: "application/json",
//     },
//   });
//   const res = await response.json();
//   let pages = Array(res.noOfPages)
//     .fill(0)
//     .map((_x, i) => ({
//       id: i.toString(),
//     }));
//   // console.log("pages", pages);
//   return pages;
// }
//
// let countries = [];
// let pages = [1, 2, 3];
// /**
//  * @param {string} index
//  */
// async function loadData(index) {
//   const [res, _status] = await fetchCountries(parseInt(index));
//   countries = res.data;
//   pages = Array(res.noOfPages)
//     .fill(0)
//     .map((_x, i) => i);
// }
//
// $: {
//   loadData($page.params["id"]);
// }
