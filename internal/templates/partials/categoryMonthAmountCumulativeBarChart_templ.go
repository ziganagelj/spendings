// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"spendings/internal/domain"
)

func graphCumulativeBar(aggregate []domain.TransactionAggCategoryMonth) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_graphCumulativeBar_a9f8`,
		Function: `function __templ_graphCumulativeBar_a9f8(aggregate){const categories = [...new Set(aggregate.map(item => item.Category))];

    const months = Array.from({ length: 12 }, (_, i) => i + 1);
    const backgroundColors = [
        'rgba(255, 99, 132, 0.2)',  // Red
        'rgba(54, 162, 235, 0.2)',  // Blue
        'rgba(255, 206, 86, 0.2)',  // Yellow
        'rgba(75, 192, 192, 0.2)',  // Green
        'rgba(153, 102, 255, 0.2)', // Purple
        'rgba(255, 159, 64, 0.2)',  // Orange
        'rgba(231, 233, 237, 0.2)', // Light Grey
        'rgba(255, 99, 255, 0.2)',  // Pink
        'rgba(99, 255, 132, 0.2)',  // Light Green
        'rgba(132, 99, 255, 0.2)'   // Light Blue
    ];

    const borderColors = [
        'rgba(255, 99, 132, 1)',  // Red
        'rgba(54, 162, 235, 1)',  // Blue
        'rgba(255, 206, 86, 1)',  // Yellow
        'rgba(75, 192, 192, 1)',  // Green
        'rgba(153, 102, 255, 1)', // Purple
        'rgba(255, 159, 64, 1)',  // Orange
        'rgba(231, 233, 237, 1)', // Light Grey
        'rgba(255, 99, 255, 1)',  // Pink
        'rgba(99, 255, 132, 1)',  // Light Green
        'rgba(132, 99, 255, 1)'   // Light Blue
    ];

    const datasets = categories.map((category, index) => {
        let cumulativeAmount = 0;
        const categoryData = Array(12).fill(0); // Initialize an array for 12 months with zeros

        for (let month = 1; month <= 12; month++) {
            aggregate.forEach(item => {
                if (item.Category === category && item.Month === month) {
                    cumulativeAmount += item.Amount;
                }
            });
            categoryData[month - 1] = cumulativeAmount; // Set cumulative amount for the month
        }

        return {
            label: category,
            data: categoryData,
            backgroundColor: backgroundColors[index % backgroundColors.length],
            borderColor: borderColors[index % borderColors.length],
            borderWidth: 1
        };
    });

    const ctx = document.getElementById('categoryMonthAmountCumulativeBarChart').getContext('2d');
    const totalMonthlySpendingChart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
            datasets: datasets
        },
        options: {
            scales: {
                x: { stacked: true },
                y: { stacked: true }
            }
        }
    });}`,
		Call:       templ.SafeScript(`__templ_graphCumulativeBar_a9f8`, aggregate),
		CallInline: templ.SafeScriptInline(`__templ_graphCumulativeBar_a9f8`, aggregate),
	}
}

func CategoryMonthAmountCumulativeBarChart(aggregate []domain.TransactionAggCategoryMonth) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, graphCumulativeBar(aggregate))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.ComponentScript = graphCumulativeBar(aggregate)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><canvas id=\"categoryMonthAmountCumulativeBarChart\"></canvas></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
