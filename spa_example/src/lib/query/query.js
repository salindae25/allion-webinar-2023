/**
 * @param {number} page
 */
export const fetchCountries = async (page) => {
  try {
    const res = await fetch(`http://localhost:8080/countries?page=${page}`, {
      method: "GET",
      headers: {
        accept: "application/json",
      },
    });
    const resJson = await res.json();
    return [resJson, "SUCESS"];
  } catch (ex) { }
  return [[], "FAILED"];
};

/**
 * @param {number} ccn
 */
export const postCountryLike = async (ccn) => {
  try {
    const res = await fetch(`http://localhost:8080/country/${ccn}/like`, {
      method: "POST",
      headers: {
        accept: "application/json",
      },
    });
    const resJson = await res.json();
    return [resJson, "SUCESS"];
  } catch (ex) { }
  return [[], "FAILED"];
};
