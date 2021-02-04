<script>
  export let location = '';
</script>

<div class="tos">

  <h2>Methodology</h2>

  <p>
    FuckTheSuits gets its data from
    <a href="http://regsho.finra.org/regsho-Index.html" target="_blank" rel="nofollow noreferrer noopener">FINRA</a>. We sum up the short volume for every company in the NYSE or NASDAQ over the previous 5 trading days, then graph that data for your amusement.
  </p>

  <p>
    Nothing on this website should be construed as financial advice.
  </p>

  <p>
    Our calculations could be catastrophically wrong, for retards we are.
  </p>

</div>

<style>
  .tos {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  p {
    font-size: 18px;
  }

  h2, h4, p {
    margin-top: 24px;
    max-width: 80%;
  }

  @media (max-width: 639px) {
    h2, h4, p {
      max-width: calc(100% - 24px);
    }
  }

</style>
