<script>
  import Footer from './Footer.svelte';

  export let location = '';
</script>

<div class="container">
  <div class="main-grid">

    <h2>Terms of Service</h2>

    <h4>Guiding philosophy: you can do anything you want with this software <em>except</em> commit or plan human rights violations.<strong></strong></h4>

    <p>
      Nothing on this website should be construed as financial advice.
    </p>

    <p>
      Our calculations could be catastrophically wrong, for retards we are.
    </p>

    <p>
      We are liable for nothing.
    </p>

    <p>
      All rights reserved, bitches.
    </p>

    <br />
    <hr />

    <p>We do not forgive.</p>

    <p>We do not forget.</p>

    <p>Expect us.</p>

    <p>United as one.</p>

    <p>Divided by zero.</p>

    <br />
    <hr />
  </div>
</div>

<style>
p {
  font-size: 18px;
}

h2, h4, p {
  margin-top: 24px;
  max-width: 100%;
}

@media (max-width: 639px) {
  h2, h4, p {
    max-width: calc(100% - 24px);
  }
}
</style>
