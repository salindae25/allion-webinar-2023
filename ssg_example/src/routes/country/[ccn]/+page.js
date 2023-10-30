export const load = async (/** @type {{ params: any; }} */ loadEvent) => {
    const { params } = loadEvent;
    const { ccn } = params;
    const response = await fetch(`http://localhost:8080/country/${ccn}`, {
        method: "GET",
        headers: {
            accept: "application/json",
        },
    });
    const res = await response.json();
    return res.data
};
