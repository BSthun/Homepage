<script lang="ts">
	import { getContext, onMount, setContext } from 'svelte';
	import { useLocation, useParams } from 'svelte-navigator';
	import Container from '../../components/layout/Container.svelte';
	import { axios, caller } from '../../utils/api';
	import SectionItem from './SectionItem.svelte';
	
	const params = useParams();
	const bind = getContext('bind');
	
	let title = 'Section';
	let state: any = {
		section: {},
	};
	
	const mount = () => {
		caller(axios.get(`/photo/entity/section/detail`, {
			params: {
				section_id: $params['section-id'],
			},
		}))
			.then((res) => {
				state = res.data;
			})
			.catch((err) => {
				$bind.openSnackbar(err.message);
			});
	};
	
	onMount(mount);
</script>

<div class="section">

</div>

<style lang="scss">

</style>