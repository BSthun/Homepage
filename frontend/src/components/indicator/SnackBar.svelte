<script lang="ts">
	import type { SnackbarComponentDev } from '@smui/snackbar'
	import Snackbar, { Actions, Label } from '@smui/snackbar'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'

	const bind: Writable<any> = getContext('bind')
	let snackbar: SnackbarComponentDev
	let message = ''

	$: bind.update((value) => ({
		...value,
		openSnackbar: (m) => {
			message = m
			snackbar.open()
		},
	}))
</script>

<div class="snackbar">
	<Snackbar bind:this={snackbar}>
		<Label>{message}</Label>
		<Actions on:click={(e) => snackbar.close()}>
			<i class="las la-times clickable" />
		</Actions>
	</Snackbar>
</div>

<style lang="scss">
	.clickable {
		cursor: pointer;
	}

	.snackbar {
		position: relative;
		z-index: 1001;
	}
</style>
