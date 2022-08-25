<script lang="ts">
	import { onMount, setContext } from 'svelte';
	import { axios, caller } from '../utils/api';
	import Container from '../components/layout/Container.svelte';
	import { useParams } from "svelte-navigator";
	
	let title = "Album"
	let state: any = {
		album: {}
	}
	
	onMount(() => {
		caller(axios.get('/photo/entity/album/detail')).then((res) => {
			state = res.data;
			setContext('state', res.data);
		});
	});
</script>

<svelte:head>
	<title>{title} - BSthun</title>
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