<script lang="ts">
	import Tooltip, { Wrapper, Title, Content, Link, RichActions } from '@smui/tooltip'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import Container from '../../components/layout/Container.svelte'

	export let state: any
	export let setYear: (any) => void
</script>

<div class="wrapper">
	<div class="heading">
		<h3>Activity Graph</h3>
		<Wrapper rich style="flex: 1">
			<span class="material-symbols-outlined icon">help</span>
			<Tooltip>
				<Content style="width: 260px">
					3 aspects of daily activity represented<br />in the color mixture of each squares.
					<div class="r-1">
						<div class="preview-square" style="background-color: dodgerblue" />
						Self
					</div>
					<div class="r-1">
						<div class="preview-square" style="background-color: forestgreen" />
						Social
					</div>
					<div class="r-1">
						<div class="preview-square" style="background-color: indianred" />
						Work
					</div>
				</Content>
			</Tooltip>
		</Wrapper>
		<select label="Activity Year" on:change={(e) => setYear(e.target.value)}>
			{#each state.local.yearOptions as year}
				<option value={year} selected={year === state.local.year}>{year}</option>
			{/each}
		</select>
	</div>
	<div class="graph">
		<div class="months">
			<span>Jan</span>
			<span>Feb</span>
			<span>Mar</span>
			<span>Apr</span>
			<span>May</span>
			<span>Jun</span>
			<span>Jul</span>
			<span>Aug</span>
			<span>Sep</span>
			<span>Oct</span>
			<span>Nov</span>
			<span>Dec</span>
		</div>
		<ul class="squares">
			{#each state.remote.days_graph as day, i}
				{#if i === 0}
					<div
						style={`background-color: #${day?.toString(16).padStart(6, '0')};
						grid-row-start: ${state.local.firstDayOffset}`}
					/>
				{:else}
					<div style={`background-color: #${day?.toString(16).padStart(6, '0')}`} />
				{/if}
			{/each}
		</ul>
	</div>
</div>

<style lang="scss">
	@import '../../styles/index';

	$gap: 4px;
	$square-size: 14px;
	$week-width: 18px;

	.wrapper {
		padding: 20px;
		margin: 20px;
		border: 1px #e1e4e8 solid;
		border-radius: 6px;
		width: calc(100% - 40px);
		max-width: 956px;
		overflow-x: scroll;
	}

	.heading {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16px;
	}

	.icon {
		margin-left: 6px;
		font-size: 18px;
		@include tweak-clickable-icon($color-grey-800);
	}

	.preview-square {
		width: 14px;
		height: 14px;
		margin-right: 6px;
	}

	.graph {
		display: inline-grid;
		grid-template-areas:
			'months'
			'squares';
		grid-template-columns: auto 1fr;
		grid-gap: $gap;
	}

	.months {
		grid-area: months;

		display: grid;
		grid-template-columns:
			calc($week-width * 4) /* Jan */
			calc($week-width * 4) /* Feb */
			calc($week-width * 5) /* Mar */
			calc($week-width * 4) /* Apr */
			calc($week-width * 4) /* May */
			calc($week-width * 5) /* Jun */
			calc($week-width * 4) /* Jul */
			calc($week-width * 4) /* Aug */
			calc($week-width * 5) /* Sep */
			calc($week-width * 4) /* Oct */
			calc($week-width * 4) /* Nov */
			calc($week-width * 5) /* Dec */;

		> span {
			font-size: 14px;
		}
	}

	.squares {
		grid-area: squares;

		display: grid;
		grid-gap: $gap;
		grid-template-rows: repeat(7, $square-size);
		grid-auto-flow: column;
		grid-auto-columns: $square-size;

		> div {
			border-radius: 2px;
		}
	}
</style>
