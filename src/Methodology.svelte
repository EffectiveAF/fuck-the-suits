<script>
  import { Link } from 'svelte-routing';

  export let location = '';
</script>

<div class="container">
  <div class="main-grid">
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

      <p>As always, don't forget to...</p>

      <Link to="/">
        <img src="/img/fuck-the-suits.svg" alt="In the end, The People will win -- always." height="65px" width="103px" style="margin-top: 24px;">
      </Link>
  </div>
</div>

<style>
p {
  font-size: 18px;
  line-height: 24px;
}

h2, p {
  margin-top: 24px;
  max-width: 100%;
}

@media (max-width: 639px) {
  .main-grid {
    margin: 0 24px;
  }
}
</style>
