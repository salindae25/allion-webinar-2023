<script>
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
        <article class="card">
          <img src={country?.flags.png} alt={country.flags.alt} />
          <header>
            <a href="/country/{country?.ccn3}">
              <h2>{country?.name.common}</h2></a
            >
          </header>
          <div class="content">
            <div>
              <span class="property">Population in country</span>:
              <span class="population">{country?.population}</span>
            </div>
            <div>
              <span class="property">Region</span>:
              <span>{country?.region}</span>
            </div>
            <div>
              <span class="property">Capital</span>:
              {#if country?.capital?.length != 0 && country?.capital}
                <span>{country?.capital[0]}</span>
              {:else}
                <span>-</span>
              {/if}
            </div>
          </div>
        </article>
      {/each}
    </div>
  {/if}
</div>
