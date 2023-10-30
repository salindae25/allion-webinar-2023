<script>
  // @ts-nocheck

  import { page } from "$app/stores";
  import { postCountryLike } from "$lib/query/query";

  /**
   * @type {{ flags: { png: any; alt: any; };
   * ccn3: any; name: { common: any; };likes:number;
   * population: any; region: any; capital: string | any[]; }
   * }
   */
  export let data;
 
  $: country = data;
</script>

<h3>Country-{$page.params["ccn"]}</h3>
{#if country}
  <h1>{country.name.common}</h1>
  <div class="container left-align">
    <div class="content">
      <img src={country.flags.png} alt={data.flags.alt} />
      <div>
        <span class="property">Population</span>:
        <span class="population">{country.population}</span>
      </div>
      <div>
        <span class="property">Region</span>:
        <span>{country.region}</span>
      </div>
      <div>
        <span class="property">Capital</span>:
        {#if country.capital != null && data.capital.length > 0}
          <span>{country.capital[0]}</span>
        {:else}
          <span>-</span>
        {/if}
      </div>
      <div>
        <span class="property">Likes</span>
        <span id="country_likes">{country.likes}</span>
      </div>
      <div>
        {#if country.flags.alt}
          <div>
            <span class="property">About Flag</span>
            <p>{country.flags.alt}</p>
          </div>
        {/if}
      </div>
    </div>
  </div>
{/if}
