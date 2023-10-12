<script>
  import { fetchCountries } from "$lib";
  import CountryCard from "$lib/component/CountryCard.svelte";
  import { onMount } from "svelte";

  let countries = [];
  /**
   * @param {string} index
   */
  async function loadData(index) {
    const [res, _status] = await fetchCountries(parseInt(index));
    countries = res.data;
  }

  onMount(async () => {
    await loadData(0);
  });
</script>

<div class="hero">
  <h3>Home</h3>
  <p>
    Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim
    labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet.
    Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum
    Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident.
    Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex
    occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat
    officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in
    Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non
    excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut
    ea consectetur et est culpa et culpa duis.
  </p>
</div>
{#if countries}
  <div class="card-container">
    {#each countries as country}
      <CountryCard data={country} />
    {/each}
  </div>

  <h4 class="align-center text-center" style="margin-top:var(--size-4)">
    <a href="/countries/0">Explore More...</a>
  </h4>
{/if}
