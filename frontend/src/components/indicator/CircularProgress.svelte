<script lang="ts">
	import CircularProgress from '@smui/circular-progress';
</script>

<div>
	<CircularProgress style="height: 32px; width: 32px;" class="colored" indeterminate fourColor />
</div>

<style global lang="scss">
	@use '@material/circular-progress/index' as circular-progress;
	@use '@material/theme/index' as theme;
	
	.colored {
		@include circular-progress.indeterminate-colors(theme.$grey-50 theme.$grey-50 theme.$grey-50 theme.$grey-50);
	}
</style>