<script>
  import { Link } from 'svelte-routing';
  import DonateButton from './DonateButton.svelte';
  import ReactivePath from './ReactivePath.svelte';
</script>

<div class="divider"></div>

<footer>
  <div class="logo-text">
    <p>Designed and Engineered in California</p>
  </div>

  <ul>
    <li>
      <Link to="/tos">Terms</Link>
    </li>
    <li>
      <Link to="/methodology">Methodology</Link>
    </li>
    <li>
      <a href="https://github.com/EffectiveAF/fuck-the-suits" target="_blank">Code</a>
    </li>
  </ul>

  <DonateButton />
</footer>

<ReactivePath />

<style>
.divider {
  width: calc(100vw - 16px);
  height: 180px;
  background-color: var(--neutral200);
}

.logo-text {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  color: var(--contentColorActive);
}

footer {
  display: flex;
  flex-direction: row;
  margin: 0 auto;
  height: 64px;
  max-width: 1250px;
}

footer > ul {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
  height: 64px;
  list-style-type: none;
}

footer ul li {
  height: 32px;
  margin: 0 24px;
  text-decoration: none;
}

footer ul > li > :global(a) {
  font-size: 16px;
}

footer > * {
  padding: 0 24px;
}

@media (max-width: 839px) {
  .divider {
    height: 64px;
  }

  .logo-text {
    justify-content: center;
  }

  footer > * {
    padding: 0 24px;
  }
}
</style>
