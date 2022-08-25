<script lang="ts">
	import Container from '../components/layout/Container.svelte';
	
	export let title = '404 Page Not Found';
	export let description = 'The page you are looking for is currently not exist on this website, try checking the URL and navigate again.';
	export let code = 'ERR_NONEXISTENT_PAGE';
</script>

<svelte:head>
	<title>Not Found - BSthun</title>
</svelte:head>

<div class="notfound">
	<Container>
		<div class="box-wrapper">
			<div class="box">
				<div class="bar"></div>
				<h1>{title}</h1>
				<p>{description}</p>
				<code>{code}</code>
			</div>
		</div>
	</Container>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.notfound {
		height: 100vh;
		background-color: $color-grey-900;
		z-index: 150;
	}
	
	.box-wrapper {
		width: 100%;
		height: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	
	.box {
		width: 450px;
		padding: 28px 20px 20px 20px;
		position: relative;
		overflow-y: hidden;
		border-radius: 8px;
		background-color: $color-grey-800;
		@include shadow-light-border;
	}
	
	.bar {
		position: absolute;
		inset: 0;
		height: 8px;
		background-color: $color-primary;
	}
	
	h1 {
		font-size: 24px;
		font-weight: 400;
		color: $color-grey-300;
		margin: 0 0 16px 0;
	}
	
	p {
		font-size: 16px;
		font-weight: 400;
		font-family: 'Assistant', Arial, serif;
		color: $color-grey-100;
		margin: 0 0 16px 0;
	}
	
	code {
		color: $color-grey-500;
	}

</style>