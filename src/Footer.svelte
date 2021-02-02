<script>
  import { Link } from 'svelte-routing';
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
  </ul>
</footer>

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
  width: 80%;
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
  justify-content: flex-end;
  align-items: center;
  width: 100%;
  height: 64px;
  list-style-type: none;
}

footer ul li {
  margin-left: 24px;
  text-decoration: none;
}

footer ul > li > :global(a) {
  font-size: 14px;
}

footer > * {
  padding: 0 24px;
}

@media (max-width: 499px) {
  .logo-text {
    width: 20%;
  }

  .logo-text > p {
    display: none;
  }

  footer > * {
    padding: 0 24px;
  }
}
</style>
