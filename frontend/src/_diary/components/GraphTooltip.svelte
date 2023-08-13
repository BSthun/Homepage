<script lang="ts">
	import Tooltip, { Wrapper, Title, Content, Link, RichActions } from '@smui/tooltip'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'

	const local: Writable<any> = getContext('local')
	const action: Writable<any> = getContext('action')

	let element: HTMLElement | null = null

	$: {
		if (element !== null && $local.activeGraph) {
			element.style.top = $local.activeGraph.event.clientY + 'px'
			element.style.left = $local.activeGraph.event.clientX + 'px'
			setTimeout(() => {
				element?.click()
			}, 0)
		}
	}
</script>

<div class="graph-tooltip">
	<Wrapper rich>
		<span class="target" bind:this={element} />
		<Tooltip persistent>
			<Title>With a Title!</Title>
			<Content>
				A persistent rich tooltip shows up when you click or press enter/space bar on an element and goes away
				when you activate it again or it loses focus. Great for informational popups on those little "i" icons.
			</Content>
		</Tooltip>
	</Wrapper>
</div>

<style lang="scss">
	.target {
		position: fixed;
	}
</style>
