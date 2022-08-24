<script lang="ts">

import Container from '../components/layout/Container.svelte';
</script>

<svelte:head>
	<title>Photo - BSthun</title>
</svelte:head>

<div class="photo">
	<Container>
		<h1 class="title">Photo</h1>
	</Container>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.photo {
		min-height: 100vh;
		background-color: $color-grey-900;
		padding-top: $size-navbar-height;
	}
	
	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
	}
</style>