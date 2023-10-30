<script>
  /**
   * @type {{ flags: { png: any; alt: any; }; ccn3: any; name: { common: any; }; population: any; region: any; capital: string | any[]; } | null}
   */
  export let data;
</script>

<article class="card">
  <img src={data?.flags.png} alt={data.flags.alt} />
  <header>
    <a href="/country/{data?.ccn3}"> <h2>{data?.name.common}</h2></a>
  </header>
  <div class="content">
    <div>
      <span class="property">Population</span>:
      <span class="population">{data?.population}</span>
    </div>
    <div>
      <span class="property">Region</span>:
      <span>{data?.region}</span>
    </div>
    <div>
      <span class="property">Capital</span>:
      {#if data?.capital?.length != 0 && data?.capital}
        <span>{data?.capital[0]}</span>
      {:else}
        <span>-</span>
      {/if}
    </div>
  </div>
</article>
