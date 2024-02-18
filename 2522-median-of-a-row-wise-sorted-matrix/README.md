<h2><a href="https://leetcode.com/problems/median-of-a-row-wise-sorted-matrix">Median of a Row Wise Sorted Matrix</a></h2> <img src='https://img.shields.io/badge/Difficulty-Medium-orange' alt='Difficulty: Medium' /><hr><p>Given an <code>m x n</code> matrix <code>grid</code> containing an <strong>odd</strong> number of integers where each row is sorted in <strong>non-decreasing</strong> order, return <em>the <strong>median</strong> of the matrix</em>.</p>

<p>You must solve the problem in less than <code>O(m * n)</code> time complexity.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> grid = [[1,1,2],[2,3,3],[1,3,4]]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The elements of the matrix in sorted order are 1,1,1,2,<u>2</u>,3,3,3,4. The median is 2.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> grid = [[1,1,3,3,4]]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The elements of the matrix in sorted order are 1,1,<u>3</u>,3,4. The median is 3.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 500</code></li>
	<li><code>m</code> and <code>n</code> are both odd.</li>
	<li><code>1 &lt;= grid[i][j] &lt;= 10<sup>6</sup></code></li>
	<li><code>grid[i]</code> is sorted in non-decreasing order.</li>
</ul>
