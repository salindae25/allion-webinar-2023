<script>
  import { page } from "$app/stores";
  import CountryCard from "$lib/component/CountryCard.svelte";
  export let data;
  $: countries = data.countries;
  $: pages = data.pages;
</script>

<h4>Countries</h4>
<div class="container">
  <nav class="page-container">
    {#each pages as pt}
      <a class="page" href="/countries/{pt}">{pt}</a>
    {/each}
  </nav>
  {#if countries && countries.length != 0}
    <div class="card-container">
      {#each countries as country}
        <CountryCard data={country} />
      {/each}
    </div>
  {/if}
</div>
